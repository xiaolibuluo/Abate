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

    public interface IHandler4250
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope4250(string Label, int Payload);

    public sealed class Handler4250 : IHandler4250
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap4251
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
    public static int Scale4252<T>(T value) where T : struct
    {
        int number = 15;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 17; }
        return number;
    }
    public interface IHandler4253
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope4253(string Label, int Payload);

    public sealed class Handler4253 : IHandler4253
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier4254(int value);

    public sealed class Source4254
    {
        public event Notifier4254 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
