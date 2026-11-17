package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked22720(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry22721 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22721 = "registry-60"
    }
}
object Registry22722 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22722 = "registry-67"
    }
}
data class Pair22723<A, B>(val first: A, val second: B)

fun combine22723(left: Int, right: Int): Pair22723<Int, Int> {
    return Pair22723(left + 74, right + 25)
}
object Registry22724 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22724 = "registry-81"
    }
}
class Counter22725 {
    var current: Int by Delegates.observable(88) { _, _, new -> println(new) }

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
sealed class Outcome22726 {
    data class Success(val payload: String) : Outcome22726()
    data object Failure : Outcome22726()
}

fun describe22726(outcome: Outcome22726): String = when (outcome) {
    is Outcome22726.Success -> outcome.payload
    Outcome22726.Failure -> "failed"
}
