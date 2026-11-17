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

    public interface IHandler13600
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope13600(string Label, int Payload);

    public sealed class Handler13600 : IHandler13600
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate13601(IEnumerable<int> source)
    {
        try
        {
            using var cursor = source.GetEnumerator();
            return await Task.FromResult(source.Where(item => item > 0).Sum());
        }
        catch (InvalidOperationException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("done");
        }
    }
    public static int Guard13602(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 45;
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
    public interface IHandler13603
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope13603(string Label, int Payload);

    public sealed class Handler13603 : IHandler13603
    {
        public void Handle(int value) { }
    }
    public interface IHandler13604
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope13604(string Label, int Payload);

    public sealed class Handler13604 : IHandler13604
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap13605
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
