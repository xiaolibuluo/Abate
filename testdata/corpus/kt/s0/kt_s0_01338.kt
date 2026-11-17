package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter11040 {
    var current: Int by Delegates.observable(98) { _, _, new -> println(new) }

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
object Registry11041 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11041 = "registry-8"
    }
}
fun <T, R> List<T>.mapChecked11042(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome11043 {
    data class Success(val payload: String) : Outcome11043()
    data object Failure : Outcome11043()
}

fun describe11043(outcome: Outcome11043): String = when (outcome) {
    is Outcome11043.Success -> outcome.payload
    Outcome11043.Failure -> "failed"
}
class Counter11044 {
    var current: Int by Delegates.observable(29) { _, _, new -> println(new) }

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
data class Pair11045<A, B>(val first: A, val second: B)

fun combine11045(left: Int, right: Int): Pair11045<Int, Int> {
    return Pair11045(left + 36, right + 45)
}
