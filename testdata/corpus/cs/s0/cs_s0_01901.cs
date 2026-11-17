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

    public delegate void Notifier17740(int value);

    public sealed class Source17740
    {
        public event Notifier17740 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard17741(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 85;
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
    public struct Point17742
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 92;
    }

    public static class Grid17742
    {
        public static Point17742 Make(int left, int right)
        {
            return new Point17742 { X = left, Y = right };
        }
    }
}
