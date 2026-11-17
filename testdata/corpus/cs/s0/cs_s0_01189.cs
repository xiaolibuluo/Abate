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

    public static int Scale11140<T>(T value) where T : struct
    {
        int number = 34;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 60; }
        return number;
    }
    public sealed class SlotMap11141
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
    public interface IHandler11142
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope11142(string Label, int Payload);

    public sealed class Handler11142 : IHandler11142
    {
        public void Handle(int value) { }
    }
    public interface IHandler11143
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope11143(string Label, int Payload);

    public sealed class Handler11143 : IHandler11143
    {
        public void Handle(int value) { }
    }
    public static int Scale11144<T>(T value) where T : struct
    {
        int number = 62;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 43; }
        return number;
    }
    public interface IHandler11145
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope11145(string Label, int Payload);

    public sealed class Handler11145 : IHandler11145
    {
        public void Handle(int value) { }
    }
}
