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

    public static int Guard26390(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 71;
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
    public struct Point26391
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 78;
    }

    public static class Grid26391
    {
        public static Point26391 Make(int left, int right)
        {
            return new Point26391 { X = left, Y = right };
        }
    }
}
