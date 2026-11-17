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

    public static int Guard24790(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 28;
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
    public static int Scale24791<T>(T value) where T : struct
    {
        int number = 35;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 49; }
        return number;
    }
    public struct Point24792
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 42;
    }

    public static class Grid24792
    {
        public static Point24792 Make(int left, int right)
        {
            return new Point24792 { X = left, Y = right };
        }
    }
    public interface IHandler24793
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope24793(string Label, int Payload);

    public sealed class Handler24793 : IHandler24793
    {
        public void Handle(int value) { }
    }
    public static int Scale24794<T>(T value) where T : struct
    {
        int number = 56;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 21; }
        return number;
    }
    public sealed class SlotMap24795
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
