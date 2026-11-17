package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter15080 {
    var current: Int by Delegates.observable(15) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
@JvmName("step15081")
fun step15081(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 22
}
fun <T, R> List<T>.mapChecked15082(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked15083(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked15084(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step15085")
fun step15085(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 50
}
data class Pair15086<A, B>(val first: A, val second: B)

fun combine15086(left: Int, right: Int): Pair15086<Int, Int> {
    return Pair15086(left + 57, right + 17)
}
class Counter15087 {
    var current: Int by Delegates.observable(64) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
sealed class Outcome15088 {
    data class Success(val payload: String) : Outcome15088()
    data object Failure : Outcome15088()
}

fun describe15088(outcome: Outcome15088): String = when (outcome) {
    is Outcome15088.Success -> outcome.payload
    Outcome15088.Failure -> "failed"
}
