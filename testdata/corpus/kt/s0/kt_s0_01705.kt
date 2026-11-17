package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry14030 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14030 = "registry-8"
    }
}
fun <T> List<T>.secondOrNull14031(): T? = if (size > 1) this[1] else null

fun render14031(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair14032<A, B>(val first: A, val second: B)

fun combine14032(left: Int, right: Int): Pair14032<Int, Int> {
    return Pair14032(left + 22, right + 31)
}
sealed class Outcome14033 {
    data class Success(val payload: String) : Outcome14033()
    data object Failure : Outcome14033()
}

fun describe14033(outcome: Outcome14033): String = when (outcome) {
    is Outcome14033.Success -> outcome.payload
    Outcome14033.Failure -> "failed"
}
@JvmName("step14034")
fun step14034(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 36
}
@JvmName("step14035")
fun step14035(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 43
}
class Counter14036 {
    var current: Int by Delegates.observable(50) { _, _, new -> println(new) }

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
@JvmName("step14037")
fun step14037(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 57
}
