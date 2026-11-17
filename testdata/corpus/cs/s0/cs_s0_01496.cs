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

    public struct Point14010
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 79;
    }

    public static class Grid14010
    {
        public static Point14010 Make(int left, int right)
        {
            return new Point14010 { X = left, Y = right };
        }
    }
    public delegate void Notifier14011(int value);

    public sealed class Source14011
    {
        public event Notifier14011 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard14012(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 93;
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
    public sealed class SlotMap14013
    {
        private readonly Dictionary<int, string> items = new Dictionary<int, string>();

        public string this[int index]
        {
            get => items.TryGetValue(index, out var value) ? value : string.Empty;
            set => items[index] = value;
        }

        public string Describe(int index) => index switch
        {
            0 => "zero",
            _ => items.ContainsKey(index) ? "present" : "absent"
        };
    }
    public static int Scale14014<T>(T value) where T : struct
    {
        int number = 10;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 39; }
        return number;
    }
}
