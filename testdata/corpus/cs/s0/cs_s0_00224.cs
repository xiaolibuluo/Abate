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

    public interface IHandler2120
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope2120(string Label, int Payload);

    public sealed class Handler2120 : IHandler2120
    {
        public void Handle(int value) { }
    }
    public interface IHandler2121
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope2121(string Label, int Payload);

    public sealed class Handler2121 : IHandler2121
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap2122
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
