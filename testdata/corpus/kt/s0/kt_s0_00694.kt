package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step5690")
fun step5690(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 30
}
fun <T, R> List<T>.mapChecked5691(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull5692(): T? = if (size > 1) this[1] else null

fun render5692(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome5693 {
    data class Success(val payload: String) : Outcome5693()
    data object Failure : Outcome5693()
}

fun describe5693(outcome: Outcome5693): String = when (outcome) {
    is Outcome5693.Success -> outcome.payload
    Outcome5693.Failure -> "failed"
}
object Registry5694 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5694 = "registry-58"
    }
}
data class Pair5695<A, B>(val first: A, val second: B)

fun combine5695(left: Int, right: Int): Pair5695<Int, Int> {
    return Pair5695(left + 65, right + 21)
}
class Counter5696 {
    var current: Int by Delegates.observable(72) { _, _, new -> println(new) }

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
