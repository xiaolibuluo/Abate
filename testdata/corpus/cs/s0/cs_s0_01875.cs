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

    public async Task<int> Aggregate17500(IEnumerable<int> source)
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
    public interface IHandler17501
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope17501(string Label, int Payload);

    public sealed class Handler17501 : IHandler17501
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap17502
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
    public delegate void Notifier17503(int value);

    public sealed class Source17503
    {
        public event Notifier17503 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point17504
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 85;
    }

    public static class Grid17504
    {
        public static Point17504 Make(int left, int right)
        {
            return new Point17504 { X = left, Y = right };
        }
    }
}
