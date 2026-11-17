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

    public interface IHandler6110
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope6110(string Label, int Payload);

    public sealed class Handler6110 : IHandler6110
    {
        public void Handle(int value) { }
    }
    public interface IHandler6111
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope6111(string Label, int Payload);

    public sealed class Handler6111 : IHandler6111
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap6112
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
