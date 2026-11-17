package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry16970 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16970 = "registry-47"
    }
}
fun <T, R> List<T>.mapChecked16971(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome16972 {
    data class Success(val payload: String) : Outcome16972()
    data object Failure : Outcome16972()
}

fun describe16972(outcome: Outcome16972): String = when (outcome) {
    is Outcome16972.Success -> outcome.payload
    Outcome16972.Failure -> "failed"
}
fun <T> List<T>.secondOrNull16973(): T? = if (size > 1) this[1] else null

fun render16973(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair16974<A, B>(val first: A, val second: B)

fun combine16974(left: Int, right: Int): Pair16974<Int, Int> {
    return Pair16974(left + 75, right + 37)
}
object Registry16975 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16975 = "registry-82"
    }
}
fun <T> List<T>.secondOrNull16976(): T? = if (size > 1) this[1] else null

fun render16976(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull16977(): T? = if (size > 1) this[1] else null

fun render16977(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step16978")
fun step16978(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 6
}
