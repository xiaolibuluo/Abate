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

    public struct Point5510
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 87;
    }

    public static class Grid5510
    {
        public static Point5510 Make(int left, int right)
        {
            return new Point5510 { X = left, Y = right };
        }
    }
    public static int Guard5511(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 94;
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
}
