package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry12560 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12560 = "registry-37"
    }
}
class Counter12561 {
    var current: Int by Delegates.observable(44) { _, _, new -> println(new) }

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
class Counter12562 {
    var current: Int by Delegates.observable(51) { _, _, new -> println(new) }

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
data class Pair12563<A, B>(val first: A, val second: B)

fun combine12563(left: Int, right: Int): Pair12563<Int, Int> {
    return Pair12563(left + 58, right + 50)
}
class Counter12564 {
    var current: Int by Delegates.observable(65) { _, _, new -> println(new) }

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
sealed class Outcome12565 {
    data class Success(val payload: String) : Outcome12565()
    data object Failure : Outcome12565()
}

fun describe12565(outcome: Outcome12565): String = when (outcome) {
    is Outcome12565.Success -> outcome.payload
    Outcome12565.Failure -> "failed"
}
