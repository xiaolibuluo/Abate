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

    public static int Scale20840<T>(T value) where T : struct
    {
        int number = 34;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 18; }
        return number;
    }
    public delegate void Notifier20841(int value);

    public sealed class Source20841
    {
        public event Notifier20841 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale20842<T>(T value) where T : struct
    {
        int number = 48;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 40; }
        return number;
    }
    public static int Guard20843(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 55;
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
    public struct Point20844
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 62;
    }

    public static class Grid20844
    {
        public static Point20844 Make(int left, int right)
        {
            return new Point20844 { X = left, Y = right };
        }
    }
    public sealed class SlotMap20845
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
}
