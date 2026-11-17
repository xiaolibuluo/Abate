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

    public async Task<int> Aggregate12040(IEnumerable<int> source)
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
    public delegate void Notifier12041(int value);

    public sealed class Source12041
    {
        public event Notifier12041 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier12042(int value);

    public sealed class Source12042
    {
        public event Notifier12042 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point12043
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 61;
    }

    public static class Grid12043
    {
        public static Point12043 Make(int left, int right)
        {
            return new Point12043 { X = left, Y = right };
        }
    }
    public interface IHandler12044
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope12044(string Label, int Payload);

    public sealed class Handler12044 : IHandler12044
    {
        public void Handle(int value) { }
    }
}
