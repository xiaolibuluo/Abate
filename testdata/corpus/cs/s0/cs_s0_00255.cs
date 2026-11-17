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

    public sealed class SlotMap2420
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
    public interface IHandler2421
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope2421(string Label, int Payload);

    public sealed class Handler2421 : IHandler2421
    {
        public void Handle(int value) { }
    }
    public static int Scale2422<T>(T value) where T : struct
    {
        int number = 61;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 17; }
        return number;
    }
    public sealed class SlotMap2423
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
    public delegate void Notifier2424(int value);

    public sealed class Source2424
    {
        public event Notifier2424 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier2425(int value);

    public sealed class Source2425
    {
        public event Notifier2425 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
