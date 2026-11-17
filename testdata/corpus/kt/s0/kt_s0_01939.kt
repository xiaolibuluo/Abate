package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair16010<A, B>(val first: A, val second: B)

fun combine16010(left: Int, right: Int): Pair16010<Int, Int> {
    return Pair16010(left + 60, right + 53)
}
fun <T> List<T>.secondOrNull16011(): T? = if (size > 1) this[1] else null

fun render16011(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked16012(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry16013 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16013 = "registry-81"
    }
}
@JvmName("step16014")
fun step16014(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 88
}
object Registry16015 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16015 = "registry-95"
    }
}
object Registry16016 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16016 = "registry-5"
    }
}
object Registry16017 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16017 = "registry-12"
    }
}
sealed class Outcome16018 {
    data class Success(val payload: String) : Outcome16018()
    data object Failure : Outcome16018()
}

fun describe16018(outcome: Outcome16018): String = when (outcome) {
    is Outcome16018.Success -> outcome.payload
    Outcome16018.Failure -> "failed"
}
