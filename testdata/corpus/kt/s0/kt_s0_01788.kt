package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair14670<A, B>(val first: A, val second: B)

fun combine14670(left: Int, right: Int): Pair14670<Int, Int> {
    return Pair14670(left + 64, right + 30)
}
object Registry14671 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14671 = "registry-71"
    }
}
fun <T> List<T>.secondOrNull14672(): T? = if (size > 1) this[1] else null

fun render14672(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
