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

    public struct Point2430
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 60;
    }

    public static class Grid2430
    {
        public static Point2430 Make(int left, int right)
        {
            return new Point2430 { X = left, Y = right };
        }
    }
    public static int Scale2431<T>(T value) where T : struct
    {
        int number = 67;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 13; }
        return number;
    }
    public interface IHandler2432
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope2432(string Label, int Payload);

    public sealed class Handler2432 : IHandler2432
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap2433
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
    public interface IHandler2434
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope2434(string Label, int Payload);

    public sealed class Handler2434 : IHandler2434
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap2435
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
