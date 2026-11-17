package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter3810 {
    var current: Int by Delegates.observable(11) { _, _, new -> println(new) }

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
object Registry3811 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3811 = "registry-18"
    }
}
fun <T> List<T>.secondOrNull3812(): T? = if (size > 1) this[1] else null

fun render3812(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step3813")
fun step3813(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 32
}
sealed class Outcome3814 {
    data class Success(val payload: String) : Outcome3814()
    data object Failure : Outcome3814()
}

fun describe3814(outcome: Outcome3814): String = when (outcome) {
    is Outcome3814.Success -> outcome.payload
    Outcome3814.Failure -> "failed"
}
@JvmName("step3815")
fun step3815(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 46
}
data class Pair3816<A, B>(val first: A, val second: B)

fun combine3816(left: Int, right: Int): Pair3816<Int, Int> {
    return Pair3816(left + 53, right + 58)
}
data class Pair3817<A, B>(val first: A, val second: B)

fun combine3817(left: Int, right: Int): Pair3817<Int, Int> {
    return Pair3817(left + 60, right + 69)
}
