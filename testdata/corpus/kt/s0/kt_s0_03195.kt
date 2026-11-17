package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step26260")
fun step26260(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 96
}
fun <T, R> List<T>.mapChecked26261(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair26262<A, B>(val first: A, val second: B)

fun combine26262(left: Int, right: Int): Pair26262<Int, Int> {
    return Pair26262(left + 13, right + 52)
}
sealed class Outcome26263 {
    data class Success(val payload: String) : Outcome26263()
    data object Failure : Outcome26263()
}

fun describe26263(outcome: Outcome26263): String = when (outcome) {
    is Outcome26263.Success -> outcome.payload
    Outcome26263.Failure -> "failed"
}
fun <T> List<T>.secondOrNull26264(): T? = if (size > 1) this[1] else null

fun render26264(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry26265 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG26265 = "registry-34"
    }
}
sealed class Outcome26266 {
    data class Success(val payload: String) : Outcome26266()
    data object Failure : Outcome26266()
}

fun describe26266(outcome: Outcome26266): String = when (outcome) {
    is Outcome26266.Success -> outcome.payload
    Outcome26266.Failure -> "failed"
}
class Counter26267 {
    var current: Int by Delegates.observable(48) { _, _, new -> println(new) }

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
object Registry26268 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG26268 = "registry-55"
    }
}
