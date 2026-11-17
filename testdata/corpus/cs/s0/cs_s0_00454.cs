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

    public static int Scale4260<T>(T value) where T : struct
    {
        int number = 14;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 63; }
        return number;
    }
    public interface IHandler4261
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope4261(string Label, int Payload);

    public sealed class Handler4261 : IHandler4261
    {
        public void Handle(int value) { }
    }
    public interface IHandler4262
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope4262(string Label, int Payload);

    public sealed class Handler4262 : IHandler4262
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier4263(int value);

    public sealed class Source4263
    {
        public event Notifier4263 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public sealed class SlotMap4264
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
