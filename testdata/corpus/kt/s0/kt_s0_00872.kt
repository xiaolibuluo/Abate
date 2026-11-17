package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair7140<A, B>(val first: A, val second: B)

fun combine7140(left: Int, right: Int): Pair7140<Int, Int> {
    return Pair7140(left + 72, right + 66)
}
object Registry7141 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7141 = "registry-79"
    }
}
fun <T> List<T>.secondOrNull7142(): T? = if (size > 1) this[1] else null

fun render7142(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step7143")
fun step7143(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 93
}
sealed class Outcome7144 {
    data class Success(val payload: String) : Outcome7144()
    data object Failure : Outcome7144()
}

fun describe7144(outcome: Outcome7144): String = when (outcome) {
    is Outcome7144.Success -> outcome.payload
    Outcome7144.Failure -> "failed"
}
class Counter7145 {
    var current: Int by Delegates.observable(10) { _, _, new -> println(new) }

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
