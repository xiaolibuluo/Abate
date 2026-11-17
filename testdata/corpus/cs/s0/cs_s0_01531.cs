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

    public sealed class SlotMap14350
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
    public delegate void Notifier14351(int value);

    public sealed class Source14351
    {
        public event Notifier14351 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler14352
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope14352(string Label, int Payload);

    public sealed class Handler14352 : IHandler14352
    {
        public void Handle(int value) { }
    }
}
