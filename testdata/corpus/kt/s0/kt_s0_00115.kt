package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter920 {
    var current: Int by Delegates.observable(37) { _, _, new -> println(new) }

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
@JvmName("step921")
fun step921(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 44
}
sealed class Outcome922 {
    data class Success(val payload: String) : Outcome922()
    data object Failure : Outcome922()
}

fun describe922(outcome: Outcome922): String = when (outcome) {
    is Outcome922.Success -> outcome.payload
    Outcome922.Failure -> "failed"
}
fun <T> List<T>.secondOrNull923(): T? = if (size > 1) this[1] else null

fun render923(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull924(): T? = if (size > 1) this[1] else null

fun render924(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry925 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG925 = "registry-72"
    }
}
fun <T> List<T>.secondOrNull926(): T? = if (size > 1) this[1] else null

fun render926(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked927(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
