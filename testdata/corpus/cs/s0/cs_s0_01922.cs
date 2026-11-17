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

    public delegate void Notifier17950(int value);

    public sealed class Source17950
    {
        public event Notifier17950 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler17951
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope17951(string Label, int Payload);

    public sealed class Handler17951 : IHandler17951
    {
        public void Handle(int value) { }
    }
    public struct Point17952
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 74;
    }

    public static class Grid17952
    {
        public static Point17952 Make(int left, int right)
        {
            return new Point17952 { X = left, Y = right };
        }
    }
    public static int Guard17953(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 81;
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
    public interface IHandler17954
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope17954(string Label, int Payload);

    public sealed class Handler17954 : IHandler17954
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate17955(IEnumerable<int> source)
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
