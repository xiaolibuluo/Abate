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

    public interface IHandler3650
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope3650(string Label, int Payload);

    public sealed class Handler3650 : IHandler3650
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap3651
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
    public delegate void Notifier3652(int value);

    public sealed class Source3652
    {
        public event Notifier3652 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
