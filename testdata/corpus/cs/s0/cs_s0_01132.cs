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

    public async Task<int> Aggregate10610(IEnumerable<int> source)
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
    public interface IHandler10611
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope10611(string Label, int Payload);

    public sealed class Handler10611 : IHandler10611
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier10612(int value);

    public sealed class Source10612
    {
        public event Notifier10612 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point10613
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 45;
    }

    public static class Grid10613
    {
        public static Point10613 Make(int left, int right)
        {
            return new Point10613 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate10614(IEnumerable<int> source)
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
