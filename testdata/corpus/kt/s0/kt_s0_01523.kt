package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair12610<A, B>(val first: A, val second: B)

fun combine12610(left: Int, right: Int): Pair12610<Int, Int> {
    return Pair12610(left + 5, right + 52)
}
data class Pair12611<A, B>(val first: A, val second: B)

fun combine12611(left: Int, right: Int): Pair12611<Int, Int> {
    return Pair12611(left + 12, right + 63)
}
class Counter12612 {
    var current: Int by Delegates.observable(19) { _, _, new -> println(new) }

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
data class Pair12613<A, B>(val first: A, val second: B)

fun combine12613(left: Int, right: Int): Pair12613<Int, Int> {
    return Pair12613(left + 26, right + 24)
}
data class Pair12614<A, B>(val first: A, val second: B)

fun combine12614(left: Int, right: Int): Pair12614<Int, Int> {
    return Pair12614(left + 33, right + 35)
}
object Registry12615 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12615 = "registry-40"
    }
}
class Counter12616 {
    var current: Int by Delegates.observable(47) { _, _, new -> println(new) }

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
data class Pair12617<A, B>(val first: A, val second: B)

fun combine12617(left: Int, right: Int): Pair12617<Int, Int> {
    return Pair12617(left + 54, right + 68)
}
sealed class Outcome12618 {
    data class Success(val payload: String) : Outcome12618()
    data object Failure : Outcome12618()
}

fun describe12618(outcome: Outcome12618): String = when (outcome) {
    is Outcome12618.Success -> outcome.payload
    Outcome12618.Failure -> "failed"
}
