package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome14010 {
    data class Success(val payload: String) : Outcome14010()
    data object Failure : Outcome14010()
}

fun describe14010(outcome: Outcome14010): String = when (outcome) {
    is Outcome14010.Success -> outcome.payload
    Outcome14010.Failure -> "failed"
}
fun <T> List<T>.secondOrNull14011(): T? = if (size > 1) this[1] else null

fun render14011(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step14012")
fun step14012(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 93
}
class Counter14013 {
    var current: Int by Delegates.observable(100) { _, _, new -> println(new) }

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
data class Pair14014<A, B>(val first: A, val second: B)

fun combine14014(left: Int, right: Int): Pair14014<Int, Int> {
    return Pair14014(left + 10, right + 39)
}
class Counter14015 {
    var current: Int by Delegates.observable(17) { _, _, new -> println(new) }

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
object Registry14016 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14016 = "registry-24"
    }
}
