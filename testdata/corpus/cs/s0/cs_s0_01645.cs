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

    public interface IHandler15390
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope15390(string Label, int Payload);

    public sealed class Handler15390 : IHandler15390
    {
        public void Handle(int value) { }
    }
    public static int Scale15391<T>(T value) where T : struct
    {
        int number = 37;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 57; }
        return number;
    }
    public struct Point15392
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 44;
    }

    public static class Grid15392
    {
        public static Point15392 Make(int left, int right)
        {
            return new Point15392 { X = left, Y = right };
        }
    }
    public static int Scale15393<T>(T value) where T : struct
    {
        int number = 51;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 18; }
        return number;
    }
    public sealed class SlotMap15394
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
