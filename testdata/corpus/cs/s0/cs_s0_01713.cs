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

    public static int Scale16000<T>(T value) where T : struct
    {
        int number = 47;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 46; }
        return number;
    }
    public delegate void Notifier16001(int value);

    public sealed class Source16001
    {
        public event Notifier16001 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard16002(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 61;
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
    public static int Guard16003(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 68;
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
    public delegate void Notifier16004(int value);

    public sealed class Source16004
    {
        public event Notifier16004 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point16005
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 82;
    }

    public static class Grid16005
    {
        public static Point16005 Make(int left, int right)
        {
            return new Point16005 { X = left, Y = right };
        }
    }
}
