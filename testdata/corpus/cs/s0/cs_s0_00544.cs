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

    public struct Point5090
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 26;
    }

    public static class Grid5090
    {
        public static Point5090 Make(int left, int right)
        {
            return new Point5090 { X = left, Y = right };
        }
    }
    public static int Guard5091(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 33;
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
    public struct Point5092
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 40;
    }

    public static class Grid5092
    {
        public static Point5092 Make(int left, int right)
        {
            return new Point5092 { X = left, Y = right };
        }
    }
}
