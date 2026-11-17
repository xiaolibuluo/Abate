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

    public interface IHandler21140
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope21140(string Label, int Payload);

    public sealed class Handler21140 : IHandler21140
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier21141(int value);

    public sealed class Source21141
    {
        public event Notifier21141 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler21142
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope21142(string Label, int Payload);

    public sealed class Handler21142 : IHandler21142
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap21143
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
