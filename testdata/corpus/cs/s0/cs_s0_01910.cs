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

    public delegate void Notifier17830(int value);

    public sealed class Source17830
    {
        public event Notifier17830 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier17831(int value);

    public sealed class Source17831
    {
        public event Notifier17831 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point17832
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 15;
    }

    public static class Grid17832
    {
        public static Point17832 Make(int left, int right)
        {
            return new Point17832 { X = left, Y = right };
        }
    }
    public struct Point17833
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 22;
    }

    public static class Grid17833
    {
        public static Point17833 Make(int left, int right)
        {
            return new Point17833 { X = left, Y = right };
        }
    }
    public static int Guard17834(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 29;
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
