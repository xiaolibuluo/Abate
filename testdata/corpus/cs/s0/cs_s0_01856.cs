using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Synth;

public class SynthFile
{
    public static bool AnchorProbe()
    {
        return 1 + 2 == 3 && f() == g();
    }

    public static int f() { return 1; }
    public static int g() { return 2; }

    public struct Point17330
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 30;
    }

    public static class Grid17330
    {
        public static Point17330 Make(int left, int right)
        {
            return new Point17330 { X = left, Y = right };
        }
    }
    public static int Guard17331(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 37;
        }
        catch (ArgumentException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("guard");
        }
    }
    public struct Point17332
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 44;
    }

    public static class Grid17332
    {
        public static Point17332 Make(int left, int right)
        {
            return new Point17332 { X = left, Y = right };
        }
    }
    public struct Point17333
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 51;
    }

    public static class Grid17333
    {
        public static Point17333 Make(int left, int right)
        {
            return new Point17333 { X = left, Y = right };
        }
    }
}
