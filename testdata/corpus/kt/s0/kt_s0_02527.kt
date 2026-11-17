package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked20790(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry20791 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG20791 = "registry-73"
    }
}
fun <T> List<T>.secondOrNull20792(): T? = if (size > 1) this[1] else null

fun render20792(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome20793 {
    data class Success(val payload: String) : Outcome20793()
    data object Failure : Outcome20793()
}

fun describe20793(outcome: Outcome20793): String = when (outcome) {
    is Outcome20793.Success -> outcome.payload
    Outcome20793.Failure -> "failed"
}
fun <T> List<T>.secondOrNull20794(): T? = if (size > 1) this[1] else null

fun render20794(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair20795<A, B>(val first: A, val second: B)

fun combine20795(left: Int, right: Int): Pair20795<Int, Int> {
    return Pair20795(left + 101, right + 38)
}
fun <T, R> List<T>.mapChecked20796(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
