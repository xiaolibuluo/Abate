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

    public delegate void Notifier23750(int value);

    public sealed class Source23750
    {
        public event Notifier23750 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale23751<T>(T value) where T : struct
    {
        int number = 41;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 53; }
        return number;
    }
    public struct Point23752
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 48;
    }

    public static class Grid23752
    {
        public static Point23752 Make(int left, int right)
        {
            return new Point23752 { X = left, Y = right };
        }
    }
    public delegate void Notifier23753(int value);

    public sealed class Source23753
    {
        public event Notifier23753 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier23754(int value);

    public sealed class Source23754
    {
        public event Notifier23754 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard23755(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 69;
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
