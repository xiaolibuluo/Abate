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

    public interface IHandler25730
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope25730(string Label, int Payload);

    public sealed class Handler25730 : IHandler25730
    {
        public void Handle(int value) { }
    }
    public interface IHandler25731
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope25731(string Label, int Payload);

    public sealed class Handler25731 : IHandler25731
    {
        public void Handle(int value) { }
    }
    public interface IHandler25732
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope25732(string Label, int Payload);

    public sealed class Handler25732 : IHandler25732
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap25733
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
    public delegate void Notifier25734(int value);

    public sealed class Source25734
    {
        public event Notifier25734 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate25735(IEnumerable<int> source)
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
}
