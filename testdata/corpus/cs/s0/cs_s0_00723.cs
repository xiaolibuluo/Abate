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

    public struct Point6780
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 89;
    }

    public static class Grid6780
    {
        public static Point6780 Make(int left, int right)
        {
            return new Point6780 { X = left, Y = right };
        }
    }
    public interface IHandler6781
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope6781(string Label, int Payload);

    public sealed class Handler6781 : IHandler6781
    {
        public void Handle(int value) { }
    }
    public interface IHandler6782
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope6782(string Label, int Payload);

    public sealed class Handler6782 : IHandler6782
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier6783(int value);

    public sealed class Source6783
    {
        public event Notifier6783 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public sealed class SlotMap6784
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
    public async Task<int> Aggregate6785(IEnumerable<int> source)
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
