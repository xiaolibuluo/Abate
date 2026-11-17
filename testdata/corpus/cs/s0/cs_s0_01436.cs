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

    public interface IHandler13450
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope13450(string Label, int Payload);

    public sealed class Handler13450 : IHandler13450
    {
        public void Handle(int value) { }
    }
    public interface IHandler13451
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope13451(string Label, int Payload);

    public sealed class Handler13451 : IHandler13451
    {
        public void Handle(int value) { }
    }
    public static int Scale13452<T>(T value) where T : struct
    {
        int number = 44;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 52; }
        return number;
    }
    public interface IHandler13453
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope13453(string Label, int Payload);

    public sealed class Handler13453 : IHandler13453
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap13454
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
