package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull13030(): T? = if (size > 1) this[1] else null

fun render13030(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome13031 {
    data class Success(val payload: String) : Outcome13031()
    data object Failure : Outcome13031()
}

fun describe13031(outcome: Outcome13031): String = when (outcome) {
    is Outcome13031.Success -> outcome.payload
    Outcome13031.Failure -> "failed"
}
@JvmName("step13032")
fun step13032(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 80
}
fun <T> List<T>.secondOrNull13033(): T? = if (size > 1) this[1] else null

fun render13033(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry13034 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13034 = "registry-94"
    }
}
data class Pair13035<A, B>(val first: A, val second: B)

fun combine13035(left: Int, right: Int): Pair13035<Int, Int> {
    return Pair13035(left + 101, right + 35)
}
fun <T, R> List<T>.mapChecked13036(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter13037 {
    var current: Int by Delegates.observable(18) { _, _, new -> println(new) }

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
