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

    public interface IHandler21220
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope21220(string Label, int Payload);

    public sealed class Handler21220 : IHandler21220
    {
        public void Handle(int value) { }
    }
    public static int Guard21221(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 50;
        }
        catch (ArgumentException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("guard");
        }
    }
    public interface IHandler21222
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope21222(string Label, int Payload);

    public sealed class Handler21222 : IHandler21222
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap21223
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
