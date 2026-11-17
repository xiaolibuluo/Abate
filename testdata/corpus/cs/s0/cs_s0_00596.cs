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

    public sealed class SlotMap5570
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
    public delegate void Notifier5571(int value);

    public sealed class Source5571
    {
        public event Notifier5571 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler5572
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope5572(string Label, int Payload);

    public sealed class Handler5572 : IHandler5572
    {
        public void Handle(int value) { }
    }
}
