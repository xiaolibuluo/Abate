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

    public static int Scale9170<T>(T value) where T : struct
    {
        int number = 92;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 23; }
        return number;
    }
    public interface IHandler9171
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope9171(string Label, int Payload);

    public sealed class Handler9171 : IHandler9171
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap9172
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
    public static int Scale9173<T>(T value) where T : struct
    {
        int number = 16;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 56; }
        return number;
    }
    public interface IHandler9174
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope9174(string Label, int Payload);

    public sealed class Handler9174 : IHandler9174
    {
        public void Handle(int value) { }
    }
}
