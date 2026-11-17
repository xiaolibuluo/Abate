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

    public sealed class SlotMap17980
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
    public static int Guard17981(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 9;
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
    public delegate void Notifier17982(int value);

    public sealed class Source17982
    {
        public event Notifier17982 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point17983
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 23;
    }

    public static class Grid17983
    {
        public static Point17983 Make(int left, int right)
        {
            return new Point17983 { X = left, Y = right };
        }
    }
    public static int Scale17984<T>(T value) where T : struct
    {
        int number = 30;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 12; }
        return number;
    }
    public struct Point17985
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 37;
    }

    public static class Grid17985
    {
        public static Point17985 Make(int left, int right)
        {
            return new Point17985 { X = left, Y = right };
        }
    }
}
