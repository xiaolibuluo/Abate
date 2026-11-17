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

    public interface IHandler5830
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope5830(string Label, int Payload);

    public sealed class Handler5830 : IHandler5830
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap5831
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
    public interface IHandler5832
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope5832(string Label, int Payload);

    public sealed class Handler5832 : IHandler5832
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate5833(IEnumerable<int> source)
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
    public interface IHandler5834
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope5834(string Label, int Payload);

    public sealed class Handler5834 : IHandler5834
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier5835(int value);

    public sealed class Source5835
    {
        public event Notifier5835 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
