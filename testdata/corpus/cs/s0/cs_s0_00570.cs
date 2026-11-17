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

    public struct Point5340
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 60;
    }

    public static class Grid5340
    {
        public static Point5340 Make(int left, int right)
        {
            return new Point5340 { X = left, Y = right };
        }
    }
    public interface IHandler5341
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope5341(string Label, int Payload);

    public sealed class Handler5341 : IHandler5341
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap5342
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
    public static int Scale5343<T>(T value) where T : struct
    {
        int number = 81;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 59; }
        return number;
    }
    public static int Guard5344(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 88;
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
    public interface IHandler5345
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope5345(string Label, int Payload);

    public sealed class Handler5345 : IHandler5345
    {
        public void Handle(int value) { }
    }
}
