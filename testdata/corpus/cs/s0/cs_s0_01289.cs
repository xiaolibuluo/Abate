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

    public static int Scale12090<T>(T value) where T : struct
    {
        int number = 8;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 54; }
        return number;
    }
    public struct Point12091
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 15;
    }

    public static class Grid12091
    {
        public static Point12091 Make(int left, int right)
        {
            return new Point12091 { X = left, Y = right };
        }
    }
    public sealed class SlotMap12092
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
    public interface IHandler12093
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope12093(string Label, int Payload);

    public sealed class Handler12093 : IHandler12093
    {
        public void Handle(int value) { }
    }
    public interface IHandler12094
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope12094(string Label, int Payload);

    public sealed class Handler12094 : IHandler12094
    {
        public void Handle(int value) { }
    }
}
