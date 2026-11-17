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

    public interface IHandler8750
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope8750(string Label, int Payload);

    public sealed class Handler8750 : IHandler8750
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier8751(int value);

    public sealed class Source8751
    {
        public event Notifier8751 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public sealed class SlotMap8752
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
    public async Task<int> Aggregate8753(IEnumerable<int> source)
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
    public delegate void Notifier8754(int value);

    public sealed class Source8754
    {
        public event Notifier8754 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point8755
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 66;
    }

    public static class Grid8755
    {
        public static Point8755 Make(int left, int right)
        {
            return new Point8755 { X = left, Y = right };
        }
    }
}
