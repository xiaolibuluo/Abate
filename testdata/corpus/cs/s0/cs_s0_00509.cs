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

    public delegate void Notifier4770(int value);

    public sealed class Source4770
    {
        public event Notifier4770 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier4771(int value);

    public sealed class Source4771
    {
        public event Notifier4771 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale4772<T>(T value) where T : struct
    {
        int number = 12;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 15; }
        return number;
    }
    public delegate void Notifier4773(int value);

    public sealed class Source4773
    {
        public event Notifier4773 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard4774(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 26;
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
    public struct Point4775
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 33;
    }

    public static class Grid4775
    {
        public static Point4775 Make(int left, int right)
        {
            return new Point4775 { X = left, Y = right };
        }
    }
}
