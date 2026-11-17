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

    public sealed class SlotMap6730
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
    public interface IHandler6731
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope6731(string Label, int Payload);

    public sealed class Handler6731 : IHandler6731
    {
        public void Handle(int value) { }
    }
    public struct Point6732
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 38;
    }

    public static class Grid6732
    {
        public static Point6732 Make(int left, int right)
        {
            return new Point6732 { X = left, Y = right };
        }
    }
    public static int Scale6733<T>(T value) where T : struct
    {
        int number = 45;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 56; }
        return number;
    }
    public interface IHandler6734
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope6734(string Label, int Payload);

    public sealed class Handler6734 : IHandler6734
    {
        public void Handle(int value) { }
    }
}
