package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair15980<A, B>(val first: A, val second: B)

fun combine15980(left: Int, right: Int): Pair15980<Int, Int> {
    return Pair15980(left + 21, right + 32)
}
data class Pair15981<A, B>(val first: A, val second: B)

fun combine15981(left: Int, right: Int): Pair15981<Int, Int> {
    return Pair15981(left + 28, right + 43)
}
object Registry15982 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15982 = "registry-35"
    }
}
sealed class Outcome15983 {
    data class Success(val payload: String) : Outcome15983()
    data object Failure : Outcome15983()
}

fun describe15983(outcome: Outcome15983): String = when (outcome) {
    is Outcome15983.Success -> outcome.payload
    Outcome15983.Failure -> "failed"
}
object Registry15984 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15984 = "registry-49"
    }
}
data class Pair15985<A, B>(val first: A, val second: B)

fun combine15985(left: Int, right: Int): Pair15985<Int, Int> {
    return Pair15985(left + 56, right + 26)
}
class Counter15986 {
    var current: Int by Delegates.observable(63) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked15987(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
