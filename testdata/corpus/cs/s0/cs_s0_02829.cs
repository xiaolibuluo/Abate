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

    public struct Point26460
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 65;
    }

    public static class Grid26460
    {
        public static Point26460 Make(int left, int right)
        {
            return new Point26460 { X = left, Y = right };
        }
    }
    public static int Guard26461(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 72;
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
