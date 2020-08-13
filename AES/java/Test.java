package mytest;

import javax.crypto.Cipher;
import javax.crypto.spec.IvParameterSpec;
import javax.crypto.spec.SecretKeySpec;
import java.util.Base64;

public class Test {
    private IvParameterSpec ivSpec;
    private SecretKeySpec keySpec;

    public Test(String srckey) {
        String key=paddingkey(srckey);
        try {
            byte[] keyBytes = key.getBytes();
            byte[] buf = new byte[16];
            for (int i = 0; i < keyBytes.length && i < buf.length; i++) {
                buf[i] = keyBytes[i];
            }
            this.keySpec = new SecretKeySpec(buf, "AES");
            this.ivSpec = new IvParameterSpec(keyBytes);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
    public String encrypt(String src) {
        try {
            byte[] origData=src.getBytes();
            Cipher cipher = Cipher.getInstance("AES/CBC/PKCS5Padding");
            cipher.init(Cipher.ENCRYPT_MODE, this.keySpec, this.ivSpec);
            byte[] re= cipher.doFinal(origData);
            return new String(Base64.getEncoder().encode(re));
        }  catch (Exception e) {
            e.printStackTrace();
        }
        return null;
    }

    public String decrypt(String src) throws Exception {

        byte[]  crypted= Base64.getDecoder().decode(src);
        Cipher cipher = Cipher.getInstance("AES/CBC/PKCS5Padding");
        cipher.init(Cipher.DECRYPT_MODE, this.keySpec, this.ivSpec);
        byte re[] =cipher.doFinal(crypted);
        return new String(re);
    }
    private static String paddingkey(String liu) {
        StringBuffer sb=new StringBuffer(liu);
        for(int i=liu.length();i<16;i++)
        {
            sb.append("0");
        }
        return sb.toString();

    }

    public static void main(String[] args) {
        Test test = new Test("0123456789abcdef");
        System.out.println(test.encrypt("{\"account\":\"test\",\"ip\":\"127.0.0.1\",\"time\":0}"));
        // outPut:  5SOZ33f22XAIqPNiiy3aFVa7+nELOf/a/eJ27oN7q4i7WXvuFPragDPsCHiE5fXK
    }

}
